<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="库存类型:" prop="goodType">
          <el-select v-model="formData.goodType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in goodTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="数量:" prop="number">
          <el-input v-model.number="formData.number" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="库存单位:" prop="unit">
          <el-select v-model="formData.unit" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in unitOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="总损失值:" prop="totalLossValue">
          <el-input v-model.number="formData.totalLossValue" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="累计值:" prop="totalValue">
          <el-input v-model.number="formData.totalValue" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总损失率:" prop="totalLossRate">
          <el-input-number v-model="formData.totalLossRate" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'InventoryDisplay'
}
</script>

<script setup>
import {
  createInventoryDisplay,
  updateInventoryDisplay,
  findInventoryDisplay
} from '@/api/inventoryDisplay'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const goodTypeOptions = ref([])
const unitOptions = ref([])
const formData = ref({
            goodType: undefined,
            number: 0,
            unit: undefined,
            totalLossValue: 0,
            totalValue: 0,
            totalLossRate: 0,
        })
// 验证规则
const rule = reactive({
               goodType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               number : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               unit : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findInventoryDisplay({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reinventDis
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    goodTypeOptions.value = await getDictFunc('goodType')
    unitOptions.value = await getDictFunc('unit')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createInventoryDisplay(formData.value)
               break
             case 'update':
               res = await updateInventoryDisplay(formData.value)
               break
             default:
               res = await createInventoryDisplay(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
