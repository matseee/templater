import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-template-modal',
  templateUrl: './template-modal.component.html',
  styleUrls: ['./template-modal.component.scss']
})
export class TemplateModalComponent implements OnInit {

  public templateName: string = '';

  constructor() { }

  ngOnInit(): void {
  }

}
