<template>
    <div>
      <warning-bar title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示" />
      <div class="gva-table-box">
        <div class="gva-btn-list">
          <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        </div>
        <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="ID"
        >
          <el-table-column type="selection" width="55" />
          <el-table-column align="left" label="接入日期" width="180">
            <template #default="scope">
              <span>{{ formatDate(scope.row.CreatedAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column align="left" label="名称" prop="type_name" width="240" />
          <el-table-column align="left" label="单位" prop="type_name" width="240" />
          <el-table-column align="left" label="操作员" prop="operator" width="240" />
          <el-table-column align="left" label="操作" min-width="260">
            <template #default="scope">
              <!-- <el-button type="primary" link icon="edit" @click="updateGoodtype(scope.row)">变更</el-button> -->
              <el-popover v-model="scope.row.visible" placement="top" width="160">
                <p>确定要删除吗？</p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                  <el-button type="primary" @click="deleteGoodType(scope.row)">确定</el-button>
                </div>
                <template #reference>
                  <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
                </template>
              </el-popover>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
          <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
          />
        </div>
      </div>
      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="客户">
        <el-form :inline="true" :model="form" label-width="80px">
            
          <el-form-item label="名称">
            <el-input v-model="form.type_name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="单位">
            <el-input v-model="form.type_name" autocomplete="off" />
          </el-form-item>
          <el-form-item label="操作员">
            <el-input v-model="form.operator" autocomplete="off" />
          </el-form-item>

        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="closeDialog">取 消</el-button>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </template>
  
  <script setup>
  import {
    createExaGoodType,
    deleteExaGoodType,
    getExaGoodType,
    getExaGoodTypeList
  } from '@/api/goodType'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { formatDate } from '@/utils/format'
  
  const form = ref({
    type_name: '',
    operator: ''
  })
  
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  
  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }
  
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }
  
  // 查询
  const getTableData = async() => {
    const table = await getExaGoodTypeList({ page: page.value, pageSize: pageSize.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  
  getTableData()
  
  const dialogFormVisible = ref(false)
  const type = ref('')
  // const updateGoodType = async(row) => {
  //   const res = await getExaGoodType({ ID: row.ID })
  //   type.value = 'update'
  //   if (res.code === 0) {
  //     form.value = res.data.customer
  //     dialogFormVisible.value = true
  //   }
  // }
  const closeDialog = () => {
    dialogFormVisible.value = false
    form.value = {
      good_type: '',
      operator: ''
    }
  }
  const deleteGoodType = async(row) => {
    row.visible = false
    const res = await deleteExaGoodType({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }
  const enterDialog = async() => {
    let res
    switch (type.value) {
      case 'create':
        res = await createExaGoodType(form.value)
        break
      default:
        res = await createExaGoodType(form.value)
        break
    }
  
    if (res.code === 0) {
      closeDialog()
      getTableData()
    }
  }
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }
  
  </script>
  
  <script>
  
  export default {
    name: 'GoodType'
  }
  </script>
  
  <style></style>
  