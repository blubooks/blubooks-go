
interface Fields {
    [key: string]: string;
}


type Nullable<T> = {
    [P in keyof T]: T[P] | null;
};

interface ResponseError  {
    fields?: Fields;
    message?: string;
}
  

export type Error = ResponseError | null
