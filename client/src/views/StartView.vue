<template>
  <div class="row">
    <div class="cols col-12">
      <h3>Mandanten</h3>        
    </div>
  </div>
  <div class="row">
    <div class="cols col-12"> 
      <div v-if="error && error!.message" class="alert alert-danger" role="alert">
          {{ error!.message }}
      </div>   
      <div v-for="client in store.clients">
       <router-link :to="{ name: 'client', params: { id: client.id } }">{{ client.title }}</router-link> 
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import type { AppError } from "@/models/app.model";
import { useAppStore } from "@/stores/app";

const loading = ref(false)
const error = ref<AppError>(null)
const store = useAppStore()

function loadData(){
  loading.value = true
  store.loadPageClients().then( 
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