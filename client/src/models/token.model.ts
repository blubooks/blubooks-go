/*
export default class Token {
    accessToken: string = ""
    refreshToken: string = ""
}
*/

export interface Token {
  email: string;
  accessToken: string;
  refreshToken: string;
}
