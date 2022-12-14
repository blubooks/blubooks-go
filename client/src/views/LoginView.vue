<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
        id="profile-img"
        src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
        class="profile-img-card"
      />
      <Form @submit="handleLogin" :validation-schema="schema">
        <div class="form-group">
          <label for="email">Email</label>
          <Field name="email" type="text" class="form-control" />
          <ErrorMessage name="email" class="error-feedback" />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <Field name="password" type="password" class="form-control" />
          <ErrorMessage name="password" class="error-feedback" />
        </div>

        <div class="form-group">
          <button class="btn btn-primary btn-block" :disabled="loading">
            <span
              v-show="loading"
              class="spinner-border spinner-border-sm"
            ></span>
            <span>Login</span>
          </button>
        </div>

        <div class="form-group">
          <div v-if="error && error!.message" class="alert alert-danger" role="alert">
            {{ error!.message }}
          </div>
        </div>
      </Form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref} from 'vue'
import { Form, Field, ErrorMessage } from "vee-validate";
import type { Error } from "@/models/app.model";
import  { genResponseError } from "@/utils/errorMessage";

import type { UserLoginForm } from "@/models/user.model";
import { useAuthStore } from "@/stores/auth";
import * as yup from "yup";
import router from "@/router";

const store = useAuthStore()
const error = ref<Error>(null)

const loading = ref(false)
const schema = yup.object({
  email: yup.string().required("email is required!").email("Email is invalid!"),
  password: yup.string().required(),
});

if (store.loggedIn) {
    router.push("/clients")
}

function handleLogin(values: any, actions: any) {
  const user =  <UserLoginForm>values
  const store = useAuthStore()
  loading.value = true
  store.login(user).then( 
    () => { 
      loading.value = false
      router.push("/");
    },
    (err: any) => { 
      loading.value = false
      error.value = genResponseError(err);
      if (error.value?.fields) {
        for (let field in error.value?.fields) {
          actions.setFieldError(field, error.value?.fields[field]);
        }
      }
    }
  ) 
}

</script>

<style scoped>
label {
  display: block;
  margin-top: 10px;
}
.card-container.card {
  max-width: 350px !important;
  padding: 40px 40px;
}
.card {
  background-color: #f7f7f7;
  padding: 20px 25px 30px;
  margin: 0 auto 25px;
  margin-top: 50px;
  -moz-border-radius: 2px;
  -webkit-border-radius: 2px;
  border-radius: 2px;
  -moz-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  -webkit-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
}
.profile-img-card {
  width: 96px;
  height: 96px;
  margin: 0 auto 10px;
  display: block;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
  border-radius: 50%;
}
.error-feedback {
  color: red;
}
</style>
