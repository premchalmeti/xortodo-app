import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { tap } from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';

import { Todo } from '../interfaces/todo';
import { TODOS } from '../mock/todos';
import { ServiceURL } from '../shared/serviceURL';

@Injectable({
  providedIn: 'root',
})
export class TodoService {
  todosAPI: string = ServiceURL.todoURL;
  deleteAPI: string = ServiceURL.todoURL;
  updateAPI: string = ServiceURL.todoURL;
  markDoneAPI: string = ServiceURL.todoMarkDoneURL;

  constructor(private http: HttpClient) {}

  fetchTodos(): Observable<Todo[]> {
    return this.http.get<Todo[]>(this.todosAPI);
  }

  deleteTodo(todo: Todo): Observable<any> {
    return this.http.delete<any>(`${this.deleteAPI}${todo.id}/`, {'observe':'response'});
  }

  updateTodo(todo: Todo): Observable<any> {
    return this.http.put<any>(`${this.updateAPI}${todo.id}/`, todo);
  }

  setDoneStatus(todo: Todo): Observable<any> {
    return this.http.post(this.markDoneAPI, { id: todo.id, done: todo.done });
  }

  createTodo(todo: Todo): Observable<any> {
    return this.http.post(this.todosAPI, todo);
  }

  getNewTodoObj(): Todo {
    return {
      title: '',
      desc: '',
      done: false,
    };
  }

  getTodoForUpdate(todo: Todo): Todo {
    return JSON.parse(JSON.stringify(todo));
  }
}
