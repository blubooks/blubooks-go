export type PageContentClient = {
    id: string
    title: string
}
export type PageContentCollection = {
    id: string
    title: string
}
export type PageContentBook = {
    id: string
    title: string
}

export type PageContentSection= {
    id: string
    title: string
    content?: string
}

export type PageContentSectionNavi= {
    id: string
    title: string
    level: number
    ids: string[]
    children?: PageContentSection[]
} | null



export type PageContent = {
    clients?: PageContentClient[]
    client?: PageContentClient
    collections?: PageContentCollection[]
    collection?: PageContentCollection
    books?: PageContentBook[]
    book?: PageContentBook
    sections?: PageContentSectionNavi[]
    section?: PageContentSection    
}

export type Page = {
    title: string,
    type: string,
    back: string, 
    content: PageContent
} 

