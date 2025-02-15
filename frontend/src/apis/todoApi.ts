import globalAxios from "./globalAxios";
import {
  TodoListResponse,
  TodoResponse,
  GetTodoRequest,
  CreateTodoRequest,
  UpdateTodoRequest,
  DeleteTodoRequest,
} from "../types/Todo";

export const getTodos = async () => {
  try {
    const response = await globalAxios.get<TodoListResponse>("/todos");
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const getTodo = async (request: GetTodoRequest) => {
  try {
    const response = await globalAxios.get<TodoResponse>(
      `/todos/${request.id}`
    );
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const createTodo = async (request: CreateTodoRequest) => {
  try {
    const response = await globalAxios.post<TodoResponse>("/todos", request);
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const updateTodo = async (request: UpdateTodoRequest) => {
  try {
    const response = await globalAxios.put<TodoResponse>(
      `/todos/${request.id}`,
      {
        title: request.title,
        content: request.content,
      }
    );
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const deleteTodo = async (request: DeleteTodoRequest) => {
  try {
    const response = await globalAxios.delete<TodoResponse>(
      `/todos/${request.id}`
    );
    return response.data;
  } catch (error) {
    console.error(error);
  }
};
