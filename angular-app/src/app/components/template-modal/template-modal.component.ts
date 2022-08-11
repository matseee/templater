import { Component, Inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { TemplateFacade } from './../../facades/template.facade';
import { Template } from './../../models/template.model';

export interface TemplateModalData {
  data: Template;
}

@Component({
  selector: 'app-template-modal',
  templateUrl: './template-modal.component.html',
  styleUrls: ['./template-modal.component.scss']
})
export class TemplateModalComponent {
  isEditMode: boolean = false;

  templateName: string = '';
  formGroup: FormGroup;

  constructor(
    @Inject(MAT_DIALOG_DATA) public myTemplate: TemplateModalData,
    public myDialogRef: MatDialogRef<TemplateModalComponent>,
    public myFormBuilder: FormBuilder,
    public myTemplateFacade: TemplateFacade,
  ) {
    this.formGroup = this.myFormBuilder.group({
      'name': ['', Validators.required],
      'template': ['', Validators.required]
    });

    if (myTemplate && myTemplate.data) {
      this.isEditMode = true;
      this.templateName = myTemplate.data.name;

      this.formGroup.setValue({
        name: myTemplate.data.name,
        template: myTemplate.data.template,
      });
    }
  }

  onCancle() {
    this.myDialogRef.close();
  }

  onSave() {
    const formGroupValue = this.formGroup.getRawValue();
    const template: Template = {
      id: this.myTemplate && this.myTemplate.data ? this.myTemplate.data.id : '',
      name: formGroupValue.name,
      template: formGroupValue.template,
    };

    if (this.isEditMode) {
      this.myTemplateFacade.updateTemplate(template);
    } else {
      this.myTemplateFacade.createTemplate(template);
    }

    this.myDialogRef.close();
  }
}
