<template>
  <div id="app">
    <DialogWrapper :transition-attrs="{ name: 'dialog' }" />
    <header>
      <div class="container">
        <nav class="navbar navbar-expand navbar-dark bg-dark">
          <router-link to="/" class="navbar-brand">home</router-link>
          <div class="navbar-nav mr-auto">
            <li class="nav-item">
              <a v-if="currentUser" to="/user" class="nav-link"
                >User</a
              >
            </li>
          </div>
          <div v-if="!currentUser" class="navbar-nav ml-auto">
            <li class="nav-item">
              <a to="/register" class="nav-link">
               Sign Up
              </a>
            </li>
            <li class="nav-item">
              <router-link to="/login" class="nav-link">
               Login
              </router-link>
            </li>
          </div>
          <div v-if="currentUser" class="navbar-nav ml-auto">
            <li class="nav-item">
              <a to="/profile" class="nav-link">
              
                {{ currentUser.email }}
              </a>
            </li>
            <li class="nav-item">
              <button class="btn btn-link nav-link" @click.prevent="logOut">
                LogOut
              </button>
            </li>
          </div>
        </nav>
      </div>
    </header>
    <main>
      <div class="container">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, computed, onBeforeUnmount } from "vue";
import { DialogWrapper } from "vue3-promise-dialog";
import { useAuthStore } from "@/stores/auth";
import EventBus from "@/common/EventBus";
import router from "./router";

const store = useAuthStore();

const currentUser = computed(() => {
  return store.user;
});

function logOut() {
  store.logout();
  router.push("/login");
}

onMounted(() => {
  EventBus.on("logout", () => {
    console.log("test");
    logOut();
  });
});
onBeforeUnmount(() => {
  EventBus.remove("logout", () => {});
});
</script>
