import type { UserLoginForm } from "@/models/user.model";
import { defineStore } from "pinia";
import AuthService from "../services/auth.service";
import TokenService from "../services/token.service";

const user = TokenService.getUser();
const initialState = user
  ? { loggedIn: true, user }
  : { loggedIn: false, user: null };

export const useAuthStore = defineStore("auth", {
  state: () => {
    return {
      user: initialState.user,
      loggedIn: initialState.loggedIn,
    };
  },
  actions: {
    login(user: UserLoginForm) {
      return AuthService.login(user).then(
        (user) => {
          this.user = user;
          return Promise.resolve(user);
        },
        (error) => {
          this.loggedIn = false;
          this.user = null;
          return Promise.reject(error);
        }
      );
    },
    logout() {
      AuthService.logout();
      this.loggedIn = false;
      this.user = null;
    },
    refreshToken(accessToken: string) {
      if (this.user == null) {
        return;
      }
      this.loggedIn = true;
      this.user = { ...this.user, accessToken: accessToken };
    },
  },
});

/*
  const count = ref(0);
  const doubleCount = computed(() => count.value * 2);
  function increment() {
    count.value++;
  }
*/
//return { count, doubleCount, increment };
//});
