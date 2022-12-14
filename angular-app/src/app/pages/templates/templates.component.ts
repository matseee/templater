import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { Subscription } from 'rxjs';
import { TemplateModalComponent } from './../../components/template-modal/template-modal.component';
import { TemplateFacade } from './../../facades/template.facade';
import { TemplateState } from './../../facades/template.state';
import { Template } from './../../models/template.model';

@Component({
  selector: 'app-templates',
  templateUrl: './templates.component.html',
  styleUrls: ['./templates.component.scss']
})
export class TemplatesPage implements OnInit, OnDestroy {
  searchFormGroup: FormGroup;

  tableDataSource: MatTableDataSource<Template>;
  displayedColumns: string[] = ['name', 'template', 'control'];

  protected subscription: Subscription;

  constructor(
    public myFormbuilder: FormBuilder,
    public myMatDialog: MatDialog,
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
    /** */
  }

  ngOnDestroy() {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
  }

  onAddTemplate() {
    this.myMatDialog.open(TemplateModalComponent, {});
  }

  onEditTemplate(template: Template) {
    this.myMatDialog.open(TemplateModalComponent, {
      data: { data: template },
    });
  }
}
