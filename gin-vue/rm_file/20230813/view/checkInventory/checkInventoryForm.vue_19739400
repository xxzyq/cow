<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="库存种类:" prop="goodType">
          <el-select v-model="formData.goodType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in goodTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="数量:" prop="newNumber">
          <el-input v-model.number="formData.newNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="损失值:" prop="lossValue">
          <el-input v-model.number="formData.lossValue" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="操作员:" prop="operator">
          <el-input v-model="formData.operator" :clearable="true" placeholder="请输入" />
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
  name: 'CheckInventory'
}
</script>

<script setup>
import {
  createCheckInventory,
  updateCheckInventory,
  findCheckInventory
} from '@/api/checkInventory'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const goodTypeOptions = ref([])
const formData = ref({
            goodType: undefined,
            newNumber: 0,
            lossValue: 0,
            operator: '',
        })
// 验证规则
const rule = reactive({
               goodType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               newNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               operator : [{
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
      const res = await findCheckInventory({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recheck
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    goodTypeOptions.value = await getDictFunc('goodType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCheckInventory(formData.value)
               break
             case 'update':
               res = await updateCheckInventory(formData.value)
               break
             default:
               res = await createCheckInventory(formData.value)
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
