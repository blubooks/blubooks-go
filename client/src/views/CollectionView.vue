<template>
  <div class="row">
    <div class="cols col-12">
      <div v-if="error && error!.message" class="alert alert-danger" role="alert">
          {{ error!.message }}
      </div>      
      <div  v-else>
          <h3>Mandant: {{ store.collection.title }}</h3>  
          <div v-for="item in store.books">
            <router-link :to="{ name: 'book', params: { id: item.id }}">{{ item.title }}</router-link> 
          </div>             
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import type { AppError } from "@/models/app.model";
import { useAppStore } from "@/stores/app";
import { useRoute } from 'vue-router'


const loading = ref(false)
const error = ref<AppError>(null)
const store = useAppStore()
const route = useRoute();


function loadData(){
  loading.value = true
  store.loadPageCollection(<string>route.params.id).then( 
    () => { 
      loading.value = false
    },
    (err: AppError) => { 
      loading.value = false
      error.value = err;
    }
  ) 
}

onMounted(() => {
  loadData()
});
</script>