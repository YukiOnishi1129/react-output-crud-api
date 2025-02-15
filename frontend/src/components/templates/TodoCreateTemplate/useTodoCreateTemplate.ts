import { useState, useCallback, ChangeEvent } from "react";
import { useNavigate } from "react-router";
import { NAVIGATION_PATH } from "../../../constants/navigation";

type UseTodoCreateTemplateParams = {
  addTodo: (title: string, content: string) => void;
};

export const useTodoCreateTemplate = ({
  addTodo,
}: UseTodoCreateTemplateParams) => {
  const navigate = useNavigate();

  /* local state */
  const [inputTitle, setInputTitle] = useState("");
  const [inputContent, setInputContent] = useState("");
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
   * Todo追加処理
   */
  const handleCreateTodo = useCallback(
    (e: ChangeEvent<HTMLFormElement>) => {
      e.preventDefault();
      if (inputTitle !== "" && inputContent !== "") {
        addTodo(inputTitle, inputContent);
        navigate(NAVIGATION_PATH.TOP);
      }
    },
    [addTodo, inputTitle, inputContent, navigate]
  );

  return {
    inputTitle,
    inputContent,
    handleChangeTitle,
    handleChangeContent,
    handleCreateTodo,
  };
};
