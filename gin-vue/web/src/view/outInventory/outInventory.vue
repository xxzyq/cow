<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
           <el-form-item label="库存类型" prop="typeName">
            <el-select v-model="searchInfo.typeName" clearable placeholder="请选择" @clear="()=>{searchInfo.typeName=undefined}">
              <el-option v-for="(item,key) in goodTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="操作员" prop="operator">
         <el-input v-model="searchInfo.operator" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="库存类型" prop="typeName" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.typeName,goodTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="数量" prop="number" width="120" />
        <el-table-column sortable align="left" label="操作员" prop="operator" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateOutInventoryFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="90px">
        <el-form-item label="库存类型:"  prop="typeName" >
          <el-select v-model="formData.typeName" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in goodTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="数量:"  prop="number" >
          <el-input v-model.number="formData.number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="操作员:"  prop="operator">
          <span>{{ nickname }}</span>
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

<script>
export default {
  name: 'OutInventory'
}
</script>

<script setup>
import {
  createOutInventory,
  deleteOutInventory,
  deleteOutInventoryByIds,
  updateOutInventory,
  findOutInventory,
  getOutInventoryList
} from '@/api/outInventory'
import {
  createInventoryDisplay,
  deleteInventoryDisplayByGoodType,
  findInventoryDisplayByGoodType
} from '@/api/inventoryDisplay'
import { getUserInfo} from '@/api/user.js'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, computed } from 'vue'

const nickname = ref('');
async function fetchNickname() {
  try {
    const userInfo = await getUserInfo();
    nickname.value = userInfo.data.userInfo.nickName;
  } catch (error) {
    console.error('Error getting nickname:', error);
  }
}
fetchNickname(); // 在组件初始化时调用函数获取昵称

// 自动化生成的字典（可能为空）以及字段
const goodTypeOptions = ref([])
const formData = ref({
        typeName: undefined,
        number: 0,
        operator: '',
        })
const inventoryDisplay = computed(() => ({
  goodType: formData.value.typeName,
  number: formData.value.number,
  unit: formData.value.unit,
  totalLossValue: 0,
  totalValue: 0,
  totalLossRate: 0.0,
}));

// 验证规则
const rule = reactive({
               typeName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               number : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               operator : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getOutInventoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    goodTypeOptions.value = await getDictFunc('goodType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteOutInventoryFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteOutInventoryByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateOutInventoryFunc = async(row) => {
    const res = await findOutInventory({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reout
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOutInventoryFunc = async (row) => {
    const res = await deleteOutInventory({ ID: row.ID })
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

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        typeName: undefined,
        number: 0,
        operator: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
  formData.value.operator=nickname
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              let displayRes
              switch (type.value) {
                case 'create':
                  // console.log(formData.value.typeName)
                  res = await createOutInventory(formData.value)
                  displayRes = await findInventoryDisplayByGoodType({ goodType: formData.value.typeName })
                  console.log(displayRes)
                  console.log(inventoryDisplay.value)
                  if (displayRes.code === 0) {
                      inventoryDisplay.value.unit = displayRes.data.reinventDis.unit
                      inventoryDisplay.value.number=displayRes.data.reinventDis.number - inventoryDisplay.value.number //更新后的数等于库存中原来的数-出库数量
                      inventoryDisplay.value.totalValue=displayRes.data.reinventDis.totalValue//总库存数不变
                      if (inventoryDisplay.value.totalLossValue==0){
                          inventoryDisplay.value.totalLossRate=0
                        }else{
                          inventoryDisplay.value.totalLossRate=inventoryDisplay.value.totalLossValue/inventoryDisplay.value.totalValue
                        }
                      displayRes = await deleteInventoryDisplayByGoodType({goodType: [formData.value.typeName]})
                      displayRes = await createInventoryDisplay(inventoryDisplay.value)
                      console.log("1234")
                    }
                    else{
                      // 报错
                      console.log("没有此库存，无法出库")
                    }
                  break
                case 'update':
                  res = await updateOutInventory(formData.value)
                  break
                default:
                  res = await createOutInventory(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
