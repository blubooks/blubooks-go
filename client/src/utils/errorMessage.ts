import type { Error } from "@/models/app.model";


export function genResponseError(error: any): Error {

    if (error.response && error.response.data && error.response.data.error) {
      return <Error>error.response.data.error;
    }
    const err = <Error>{};
    err!.message = error.message || error.toString();
    return err;
  }
  