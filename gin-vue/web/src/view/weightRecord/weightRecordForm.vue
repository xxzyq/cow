<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="顺序:" prop="cowNumber">
          <el-input v-model.number="formData.cowNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="毛重:" prop="routhWeight">
          <el-input v-model.number="formData.routhWeight" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="毛重时间:" prop="roughTime">
          <el-input v-model="formData.roughTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="净重:" prop="netWeight">
          <el-input v-model.number="formData.netWeight" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="净重时间:" prop="netTime">
          <el-input v-model="formData.netTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="操作员:" prop="operator">
          <el-input v-model="formData.operator" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="调整值:" prop="changeNumber">
          <el-input v-model.number="formData.changeNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片1:" prop="photo1">
          <SelectImage v-model="formData.photo1" />
       </el-form-item>
        <el-form-item label="图片组2:" prop="photo2">
           <SelectImage v-model="formData.photo2" multiple />
       </el-form-item>
        <el-form-item label="图片组3:" prop="photo3">
           <SelectImage v-model="formData.photo3" multiple />
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
  name: 'WeightRecord'
}
</script>

<script setup>
import {
  createWeightRecord,
  updateWeightRecord,
  findWeightRecord
} from '@/api/weightRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import SelectImage from '@/components/selectImage/selectImage.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            cowNumber: 0,
            routhWeight: 0,
            roughTime: '',
            netWeight: 0,
            netTime: '',
            operator: '',
            changeNumber: 0,
            remark: '',
            photo1: "",
            photo2: [],
            photo3: [],
        })
// 验证规则
const rule = reactive({
               cowNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               routhWeight : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               netWeight : [{
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
      const res = await findWeightRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reweight
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createWeightRecord(formData.value)
               break
             case 'update':
               res = await updateWeightRecord(formData.value)
               break
             default:
               res = await createWeightRecord(formData.value)
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
