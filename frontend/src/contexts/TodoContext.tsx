import { createContext, FC, ReactNode } from "react";
import { TodoType } from "../types/Todo";
import { useTodo } from "../hooks/useTodo";

type TodoContextType = {
  originTodoList: Array<TodoType>;
  addTodo: (title: string, content: string) => void;
  updateTodo: (id: number, title: string, content: string) => void;
  deleteTodo: (targetId: number, targetTitle: string) => void;
};

/**
 * TodoContext
 */
const TodoContext = createContext<TodoContextType>({
  originTodoList: [],
  addTodo: () => {},
  updateTodo: () => {},
  deleteTodo: () => {},
});

// TodoContextをエクスポート
export { TodoContext };

type TodoProviderProps = {
  children: ReactNode;
};

/**
 * TodoProvider
 * @param children
 * @constructor
 */
export const TodoProvider: FC<TodoProviderProps> = ({ children }) => {
  // カスタムフックから状態とロジックを呼び出してコンテキストプロバイダーにあてがう
  const { originTodoList, addTodo, updateTodo, deleteTodo } = useTodo();

  return (
    <TodoContext.Provider
      value={{
        originTodoList,
        addTodo,
        updateTodo,
        deleteTodo,
      }}
    >
      {children}
    </TodoContext.Provider>
  );
};
