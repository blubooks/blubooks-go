<template>
    <div class="row">
      <div class="cols col-12">
        <h3>Mandant: {{ store.book.title }}</h3>    
      </div>  


      <div  v-if="!isOpenForm" class="cols col-2  sticky-nav border-end">
        <router-link v-if="store.book && store.book.id" :to="{ name: 'book', params: { id: store.book.id }}">{{ store.book.title }}</router-link>
        <template v-if="route.name =='book'">
        <ul class="nav flex-column" v-for="item in store.sections" >
          <li class="nav-item"><router-link  class="nav-link" :to="{ name: 'section', params: { id: item.id }}">{{ item.title }}</router-link></li>
        </ul>
        </template>
        <template v-else>
        <ul class="nav flex-column" v-if="currentNavi">
          <NaviNode v-for="node in store.sections" :key="node.id" :node="node" :current="currentNavi"  /> 
        </ul>
        </template>
      </div>
      <div class="cols col-12" v-if="isOpenForm">
        <div v-if="error && error!.message" class="alert alert-danger" role="alert">
            {{ error!.message }}
        </div>     
        <Form @submit="saveForm" :validation-schema="schema">
          <div>
            <div class="form-group mb-4">
              <Field name="title" type="text" class="form-control"  v-model="editContent.title" />
              <ErrorMessage name="title" class="error-feedback" />
            </div>
            <div class="form-group mb-4">
              <Field name="content" as="textarea" rows="20" class="form-control" v-model="editContent.content" />
              <ErrorMessage name="content" class="error-feedback" />
            </div>
            <div class="form-group mb-2">
              <button class="btn btn-primary btn-block" :disabled="loading">
                <span
                  v-show="loading"
                  class="spinner-border spinner-border-sm"
                ></span>
                <span>Save</span>
              </button>
            </div>
          </div>
        </Form>      

      </div>  
      <div v-else  class="cols col-10">
        <div v-if="error && error!.message" class="alert alert-danger" role="alert">
          {{ error!.message }}
        </div>        
        <div>
          <button type="button" class="btn btn-primary" @click="openEdit()">edit</button> <button type="button" class="btn btn-primary" @click="openForm(false)">add</button>
          <div v-if="route.name =='book'">
            <div  v-html="store.book.title"></div>        
          </div>
          <div v-else>
            <h1>{{ store.section.title }}</h1>
            <div  v-html="store.section.content"></div>        
          </div>
        </div>
      </div>
    </div>

</template>

<style>
    .active {
      font-weight: bold;
    }

    .sticky-nav{        
        position: sticky;
        top:  0px;
        overflow-y: auto;
        overflow-x: hidden;
        height: calc(100vh - 98px );
    }
</style>

<script lang="ts" setup>
import { ref, onMounted, watch } from 'vue'
import type { AppError } from "@/models/app.model";
import { useAppStore } from "@/stores/app";
import { useRoute } from 'vue-router'
import NaviNode from '@/components/helper/NaviNode.vue'
import type { Page, PageContentClient, PageContentCollection, PageContentBook, PageContentSection, PageContentSectionNavi } from "@/models/page.model";
import ClientService from "@/services/client.service";
import { genResponseError } from "@/utils/errorMessage";
import * as yup from "yup";
import { Form, Field, ErrorMessage } from "vee-validate";

const loading = ref(false)
const editContent = ref({})
const isEdit = ref(false)
const isOpenForm = ref(false)

const currentNavi = ref<PageContentSectionNavi>(null)

const error = ref<AppError>(null)
const store = useAppStore()
const route = useRoute();

const schema = yup.object({
  title: yup.string().required(),
  content: yup.string(),
});

function findObject(collection: any, key: any, value: any): any {
    for (const o of collection) {
      for (let k in o) {
        if (k === key && o[k] === value) {
          return o
        }
        if (Array.isArray(o[k])) {
          const _o = findObject(o[k], key, value)
          if (_o) {
            return _o
          }
        }
      }
    }
  }

function loadData(){
  loading.value = true
  isOpenForm.value = false;
  if (route.name == "book") {
      store.loadPageBook(<string>route.params.id).then( 
      () => { 
        currentNavi.value = findObject(store.sections,"id", route.params.id);
        loading.value = false
      },
      (err: AppError) => { 
        loading.value = false
        error.value = err;
      }
    ) 
    return
  }

  store.loadPageSection(<string>route.params.id).then( 
    () => { 
      currentNavi.value = findObject(store.sections,"id", route.params.id);
      loading.value = false
    },
    (err: AppError) => { 
      loading.value = false
      error.value = err;
    }
  ) 
}

function openForm(pIsEdit: boolean){
      isEdit.value = pIsEdit;
      isOpenForm.value = true
      editContent.value = {}
}
function saveForm(value: any, actions: any){
  if (isEdit.value) {
    ClientService.putSection(<string>route.params.id,  editContent.value).then(
      (response) => {
        if (response.status === 202) {
          isOpenForm.value = false
          editContent.value = {}
          loadData();
          return;
        }
        error.value = genResponseError("Konte nicht gespeichert werden");
  
      },
      (err) => {
        error.value = genResponseError(err);
        if (error.value?.fields) {
          for (let field in error.value?.fields) {
            actions.setFieldError(field, error.value?.fields[field]);
          }
        }
      }
    );      
  }else {
    ClientService.postSection(<string>route.params.id,  editContent.value).then(
      (response) => {
        if (response.status === 201) {
          isOpenForm.value = false
          editContent.value = {}
          loadData();
          return;
        }
        editContent.value = genResponseError("Konte nicht gespeichert werden");
  
      },
      (err) => {
        error.value = genResponseError(err);
        if (error.value?.fields) {
          for (let field in error.value?.fields) {
            actions.setFieldError(field, error.value?.fields[field]);
          }
        }
      }
    ); 
  } 
}


function openEdit(){
  ClientService.getSection(<string>route.params.id).then(
    (response) => {
      openForm(true)
      editContent.value = response.data;
    },
    (error) => {
      error.value = genResponseError(error)
    }
  );
}

watch(() => route.params.id, loadData)



onMounted(() => {
  loadData()
});

/*
import ClientService from "@/services/client.service";
import EventBus from "@/common/EventBus";
import NaviNode from '@/components/helper/NaviNode.vue'


export default {
  name: "SectionView",
  components: {
    NaviNode
  },  
  data() {
    return {
      new: false,
      hasError: false,
      hasLoaded: false,
      data: "",
      editContent: null,
      formError: null,
    };
  },
  watch: {
    '$route.params.id': function(val, oldVal){ // Same
    //'$route.path': function(val, oldVal){
      console.log("changed")
      this.loadData();
    }
  },
  mounted() {
    this.loadData();

  },
  methods: {

    findObject(collection, key, value) {
      for (const o of collection) {
        for (const [k, v] of Object.entries(o)) {
          if (k === key && v === value) {
            return o
          }
          if (Array.isArray(v)) {
            const _o = this.findObject(v, key, value)
            if (_o) {
              return _o
            }
          }
        }
      }
    },

    saveClick(){
      if (this.new) {
        ClientService.postSection(this.$route.params.id,  this.editContent).then(
        (response) => {
          if (response.status === 201) {
            this.loadData();
          }
     
        },
        (error) => {
          console.log(error.response.data.error)
          if (error.response && error.response.data && error.response.data.error) {
            this.formError = error.response.data.error.message
            return;
          }
          this.formError =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();
        }
      );        
        return
      }
      ClientService.putSection(this.$route.params.id,  this.editContent).then(
        (response) => {
          if (response.status === 202) {
            this.loadData();
          }
     
        },
        (error) => {
          console.log(error.response.data.error)
          if (error.response && error.response.data && error.response.data.error) {
            this.formError = error.response.data.error.message
            return;
          }
          this.formError =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();
        }
      );
    },    
    editClick(){
      ClientService.getSection(this.$route.params.id).then(
        (response) => {
          this.editContent = response.data;
        },
        (error) => {
          this.editContent =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();
        }
      );
    }, 
    addClick(){
      this.new = true;
      this.editContent = {}
    },    
    loadData() {
      this.hasLoaded = false;
      this.false = true;

      ClientService.getPageSection(this.$route.params.id).then(
        (response) => {
          this.currentNavi = this.findObject(response.data.content.sections,"id", this.$route.params.id);

          this.editContent = null;
          this.hasLoaded = true;
          this.data = response.data;
        },
        (error) => {
          this.hasError = true;
          this.hasLoaded = true;
          this.data =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();


          if (error.response && error.response.status === 403) {
            EventBus.dispatch("logout");
            return;
          }
          if (error.response && error.response.status === 404) {

          }
          console.log(this.data)

        }
      );
    },
  }  
};
*/
</script>
<style>
.error-feedback {
  color: red;
}
</style>