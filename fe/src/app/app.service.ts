import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Database, Table, ColumnForm } from './app.model';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AppService {

  constructor(private httpClient: HttpClient) { }

  getDatabases(): Observable<Database[]> {
    return this.httpClient.get<Database[]>('http://localhost:3000/db');
  }

  getTables(name: string): Observable<Table[]> {
    return this.httpClient.get<Table[]>(`http://localhost:3000/db/${name}/tables`);
  }

  getTableForms(dbName, tableName: string): Observable<ColumnForm> {
    return this.httpClient.get<ColumnForm>(`http://localhost:3000/db/${dbName}/tables/${tableName}`);
  }
}
