import type { Token } from "@/models/token.model";

class TokenService {
  getUser(): Token | null {
    const userString = localStorage.getItem("user");
    if (userString != null) {
      return JSON.parse(userString);
    }
    return null;
  }

  getLocalRefreshToken() {
    const user = this.getUser();
    return user?.refreshToken;
  }

  getLocalAccessToken() {
    const user = this.getUser();
    return user?.accessToken;
  }

  updateLocalAccessToken(token: string) {
    const user = this.getUser();
    if (user == null) {
      return;
    }
    user.accessToken = token;
    localStorage.setItem("user", JSON.stringify(user));
  }

  setUser(user: Token) {
    console.log(JSON.stringify(user));
    localStorage.setItem("user", JSON.stringify(user));
  }

  removeUser() {
    localStorage.removeItem("user");
  }
}

export default new TokenService();
