import api from './api';

class ClientService {
 
  getClients() {
    return api.get('/page/clients');
  }
  getClient(id: string) {
    return api.get('/page/client/' + id);
  }
  getCollection(id: string) {
    return api.get('/page/collection/' + id);
  }
  getPageSection(id: string) {
    return api.get('/page/section/' + id);
  }  
  getSection(id: string) {
    return api.get('/section/' + id);
  }    
  putSection(id: string, params: any) {
    return api.put('/section/' + id, params);
  }      
  postSection(id: string, params: any) {
    return api.post('/section/' + id, params);
  }      
  getBook(id: string) {
    return api.get('/page/book/' + id);
  }  
  getClientCollections(id: string) {
    return api.get('/client/' + id + '/collections');
  }

  getUserBoard() {
    return api.get('/test/user');
  }

  getModeratorBoard() {
    return api.get('/test/mod');
  }

  getAdminBoard() {
    return api.get('/test/admin');
  }
}

export default new ClientService();