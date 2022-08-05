import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatTableDataSource } from '@angular/material/table';
import { Subscription } from 'rxjs';
import { TemplateFacade } from './../../facades/template.facade';
import { TemplateState } from './../../facades/template.state';
import { Template } from './../../models/template.model';
import { LogService } from './../../services/log.service';

@Component({
  selector: 'app-templates',
  templateUrl: './templates.component.html',
  styleUrls: ['./templates.component.scss']
})
export class TemplatesPage implements OnInit, OnDestroy {
  searchFormGroup: FormGroup;

  tableDataSource: MatTableDataSource<Template>;
  displayedColumns: string[] = ['description', 'template', 'variables', 'control'];

  protected subscription: Subscription;

  constructor(
    public myFormbuilder: FormBuilder,
    public myLogger: LogService,
    public myTemplateFacade: TemplateFacade,
  ) {
    this.searchFormGroup = this.myFormbuilder.group({
      search: ['']
    });

    this.tableDataSource = new MatTableDataSource<Template>([]);

    this.subscription = this.myTemplateFacade.state$
      .subscribe((state: TemplateState) => {
        this.tableDataSource.data = state.templates;
      });

    this.searchFormGroup.valueChanges
      .subscribe(value => console.log(value));
  }

  ngOnInit() {
    this.myLogger.debug('TemplatePage.onInit()');
  }

  ngOnDestroy() {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
  }

  onAddTemplate() {
    this.myLogger.debug('TemplatePage.onAddTemplate()');
  }

}
