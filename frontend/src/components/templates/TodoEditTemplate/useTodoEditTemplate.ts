import { useMemo, useState, useCallback, ChangeEvent } from "react";
import { useParams, useNavigate } from "react-router";
import { NAVIGATION_PATH } from "../../../constants/navigation";
import { TodoType } from "../../../types/Todo";

type UseTodoEditTemplateParams = {
  originTodoList: Array<TodoType>;
  updateTodo: (id: number, title: string, content: string) => void;
};

export const useTodoEditTemplate = ({
  originTodoList,
  updateTodo,
}: UseTodoEditTemplateParams) => {
  const navigate = useNavigate();

  const { id } = useParams();

  const todo = useMemo(
    () => originTodoList.find((todo) => String(todo.id) === id),
    [id, originTodoList]
  );
  /* local state */
  const [inputTitle, setInputTitle] = useState(todo?.title || "");
  const [inputContent, setInputContent] = useState(todo?.content || "");

  /**
   * 「title」変更処理
   */
  const handleChangeTitle = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => setInputTitle(e.target.value),
    []
  );

  /**
   * 「content」変更処理
   */
  const handleChangeContent = useCallback(
    (e: ChangeEvent<HTMLTextAreaElement>) => setInputContent(e.target.value),
    []
  );

  /**
   * Todo更新処理
   */
  const handleUpdateTodo = useCallback(
    (e: ChangeEvent<HTMLFormElement>) => {
      e.preventDefault();
      if (!!todo?.id && inputTitle !== "" && inputContent !== "") {
        updateTodo(todo.id, inputTitle, inputContent);
        navigate(NAVIGATION_PATH.TOP);
      }
    },
    [navigate, todo?.id, inputTitle, inputContent, updateTodo]
  );

  return {
    todo,
    inputTitle,
    inputContent,
    handleChangeTitle,
    handleChangeContent,
    handleUpdateTodo,
  };
};
