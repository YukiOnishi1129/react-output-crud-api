import { ApiResponse } from "./ApiResponse";

export type TodoType = {
  id: number;
  title: string;
  content?: string;
  createdAt: string;
  updatedAt: string;
};

export type TodoListType = {
  todos: Array<TodoType>;
  total: number;
};

export type TodoListResponse = ApiResponse<TodoListType>;

export type TodoResponse = ApiResponse<TodoType>;

export type GetTodoRequest = {
  id: string;
};

export type CreateTodoRequest = {
  title: string;
  content?: string;
};

export type UpdateTodoRequest = {
  id: string;
  title: string;
  content?: string;
};

export type DeleteTodoRequest = {
  id: string;
};
