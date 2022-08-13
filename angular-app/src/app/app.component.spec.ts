import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MaterialModule } from './material.module';

describe('AppComponent', () => {
  let fixture: ComponentFixture<AppComponent>;
  let domElement: HTMLElement;
  let appComponent: AppComponent;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppRoutingModule,
        BrowserAnimationsModule,
        ReactiveFormsModule,
        HttpClientModule,
        MaterialModule,
      ],
      declarations: [
        AppComponent
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(AppComponent);
    domElement = fixture.debugElement.nativeElement;
    appComponent = fixture.componentInstance;
  });

  it('should create the app', () => {
    expect(appComponent).toBeTruthy();
  });

  it('should have the title "templater"', () => {
    const title = domElement.querySelector('.title');

    expect(title).toBeTruthy();

    if (title) {
      expect(title.innerHTML).toBe('templater');
    }
  });

  it('should have a sidenav with the items "Templates" and "Settings"', () => {
    const sidenavContainer = domElement.querySelector('.sidenav-container');
    expect(sidenavContainer).toBeTruthy();

    const matListItems = domElement.querySelectorAll('.menu-item');
    expect(matListItems[0].innerHTML).toBe('Templates');
    expect(matListItems[1].innerHTML).toBe('Settings');
  });
});
