import { Component, OnInit, Input } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { Router } from '@angular/router';
import { catchError, map, tap } from 'rxjs/operators';

import { Todo } from 'src/app/interfaces/todo';
import { AppConstants } from 'src/app/shared/appConstants';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss'],
})
export class NavbarComponent implements OnInit {
  title: string = AppConstants.TITLE;

  constructor(
    private authService: AuthService,
    private router: Router,
  ) {}

  ngOnInit(): void {
  }

  logout(): void {
    this.authService.logout().subscribe((response) => {
      if (response.ok) {
        this.router.navigate(['login']);
      }
    });
  }

}
