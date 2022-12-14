
interface Fields {
    [key: string]: string;
}
interface ResponseError  {
    fields?: Fields;
    message?: string;
}


export type AppError = ResponseError | null


export type ContentClients = {
    id: string,
    title: string
}
export type ContentClient = {
    id: string,
    title: string,
    content?: string
}

export type ContentCollections = {
    id: string,
    title: string,
}
    

    







    
