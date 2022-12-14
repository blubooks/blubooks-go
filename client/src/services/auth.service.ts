import api from "./api";
import TokenService from "./token.service";
import type { UserLoginForm } from "@/models/user.model";

class AuthService {
  login(user: UserLoginForm) {
    //console.log(user)
    return api.post("/auth/login", user).then((response) => {
      if (response.data.accessToken) {
        TokenService.setUser(response.data);
      }

      return response.data;
    });
  }

  logout() {
    TokenService.removeUser();
  }

  register(user: UserLoginForm) {
    return api.post("/auth/register", user);
  }
}

export default new AuthService();
