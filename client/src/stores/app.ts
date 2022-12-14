
import { defineStore } from "pinia";
import ClientService from "@/services/client.service";
import { genResponseError } from "@/utils/errorMessage";
import type { Page, PageContentClient, PageContentCollection, PageContentBook, PageContentSection, PageContentSectionNavi } from "@/models/page.model";


export const useAppStore = defineStore("app", {
  state: () => {
    return {
      clients: [] as PageContentClient[],
      client: {} as PageContentClient,
      collection: {} as PageContentCollection,
      collections: [] as PageContentCollection[],
      book: {} as PageContentBook,
      books: [] as PageContentBook[],
      section: {} as PageContentSection,
      sections: [] as PageContentSectionNavi[],
    };
  },
  actions: {
    parsingResponse(page: Page): boolean {
      let parsed = false;
      if (page.content.clients) {
        this.clients = page.content.clients;
        parsed = true;
      }
      if (page.content.client) {
        this.client = page.content.client;
        parsed = true;
      }
      if (page.content.collection) {
        this.collection = page.content.collection;
        parsed = true;
      }
      if (page.content.collections) {
        this.collections = page.content.collections;
        parsed = true;
      }
      if (page.content.section) {
        this.section = page.content.section;
        parsed = true;
      }
      if (page.content.sections) {
        this.sections = page.content.sections;
        parsed = true;
      }
      if (page.content.book) {
        this.book = page.content.book;
        parsed = true;
      }
      if (page.content.books) {
        this.books = page.content.books;
        parsed = true;
      }
      return parsed

    },
    loadPageClients() {
      return new Promise((resolve, reject) => {
        return ClientService.getClients().then(
          (response) => {
            try {
              if (this.parsingResponse(<Page>response.data)) {
                resolve(true)
              } else {
                reject(genResponseError("Parsing Error"))
              }
            }
            catch (err) {
              reject(genResponseError(err))
            }
          },
          (error) => {
            reject(genResponseError(error))
          }
        );

      });
    },
    loadPageClient(id: string) {
      return new Promise((resolve, reject) => {
        return ClientService.getClient(id).then(
          (response) => {
            try {
              if (this.parsingResponse(<Page>response.data)) {
                resolve(true)
              } else {
                reject(genResponseError("Parsing Error"))
              }
            }
            catch (err) {
              reject(genResponseError(err))
            }
          },
          (error) => {
            reject(genResponseError(error))
          }
        );
      });
    },
    loadPageCollection(id: string) {
      return new Promise((resolve, reject) => {
        return ClientService.getCollection(id).then(
          (response) => {
            try {
              if (this.parsingResponse(<Page>response.data)) {
                resolve(true)
              } else {
                reject(genResponseError("Parsing Error"))
              }
            }
            catch (err) {
              reject(genResponseError(err))
            }
          },
          (error) => {
            reject(genResponseError(error))
          }
        );
      });
    },
    loadPageSection(id: string) {
      return new Promise((resolve, reject) => {
        return ClientService.getPageSection(id).then(
          (response) => {
            try {
              if (this.parsingResponse(<Page>response.data)) {
                resolve(true)
              } else {
                reject(genResponseError("Parsing Error"))
              }
            }
            catch (err) {
              reject(genResponseError(err))
            }
          },
          (error) => {
            reject(genResponseError(error))
          }
        );
      });
    },
    loadSection(id: string) {
      return new Promise((resolve, reject) => {
        return ClientService.getSection(id).then(
          (response) => {
            try {
              if (this.parsingResponse(<Page>response.data)) {
                resolve(true)
              } else {
                reject(genResponseError("Parsing Error"))
              }
            }
            catch (err) {
              reject(genResponseError(err))
            }
          },
          (error) => {
            reject(genResponseError(error))
          }
        );
      });
    },
    loadPageBook(id: string) {
      return new Promise((resolve, reject) => {
        return ClientService.getBook(id).then(
          (response) => {
            try {
              if (this.parsingResponse(<Page>response.data)) {
                resolve(true)
              } else {
                reject(genResponseError("Parsing Error"))
              }
            }
            catch (err) {
              reject(genResponseError(err))
            }
          },
          (error) => {
            reject(genResponseError(error))
          }
        );
      });
    },
  },
});