import { useMemo, useCallback } from "react";
import { useParams, useNavigate } from "react-router";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { NAVIGATION_PATH } from "../../../constants/navigation";
import { TodoType } from "../../../types/Todo";

const schema = z.object({
  title: z
    .string()
    .min(1, "タイトルは必須です。")
    .max(10, "10文字以内で入力してください。"),
  content: z.string().optional(),
});

type UseTodoEditTemplateParams = {
  originTodoList: Array<TodoType>;
  updateTodo: (id: number, title: string, content?: string) => void;
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

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: { title: todo?.title, content: todo?.content },
  });

  const handleEditSubmit = handleSubmit(
    useCallback(
      (values: z.infer<typeof schema>) => {
        if (!todo) return;
        updateTodo(todo.id, values.title, values.content);
        navigate(NAVIGATION_PATH.TOP);
      },
      [updateTodo, navigate, todo]
    )
  );

  return {
    todo,
    control,
    errors,
    handleEditSubmit,
  };
};
