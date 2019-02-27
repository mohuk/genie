import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';

import { AppService } from './app.service';
import { Database, Table, ColumnForm, FormlyTemplate } from './app.model';
import { MatRadioChange, MatSelectionList } from '@angular/material';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [AppService]
})
export class AppComponent implements OnInit {
  databases: Database[];
  tables: Table[];
  columnForms: ColumnForm = {formlyTemplate: []};

  selectedDb: string;
  selectedColumns: any[] = [];
  jsonObj: any;

  constructor(private appService: AppService) { }

  ngOnInit() {
    this.getDatabases();
  }

  getDatabases() {
    this.appService.getDatabases()
    .subscribe(db => this.databases = db);
  }

  getTables(dbName) {
    this.appService.getTables(dbName)
    .subscribe(tables => this.tables = tables);
  }

  onDatabaseChange(event: MatRadioChange) {
    this.selectedDb = event.value;
    this.appService.getTables(this.selectedDb)
      .subscribe(tables => this.tables = tables);
  }

  onTableChange(event: MatRadioChange) {
    const tableName = event.value;
    this.appService.getTableForms(this.selectedDb, tableName)
      .subscribe(columnForms => this.columnForms = columnForms);
  }

  onSelectColumns(selectedColumns: MatSelectionList) {
    this.selectedColumns = selectedColumns.selectedOptions
    .selected.map(item => item.value);
    this.jsonObj = JSON.parse(JSON.stringify(this.selectedColumns));
  }
}
