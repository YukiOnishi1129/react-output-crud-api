import { useCallback, useEffect } from "react";
import { useParams, useNavigate } from "react-router";
import { Controller, useForm } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

import { BaseLayout } from "../../organisms";
import { InputFormSection, TextAreaSection } from "../../molecules";
import { CommonButton } from "../../atoms";
import { getTodo, updateTodo } from "../../../apis/todo";
import { NAVIGATION_PATH } from "../../../constants/navigation";

import styles from "./style.module.css";

const schema = z.object({
  title: z
    .string()
    .min(1, "タイトルは必須です。")
    .max(10, "10文字以内で入力してください。"),
  content: z.string().optional(),
});

export const TodoEditTemplate = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  const {
    control,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
  });

  const fetchTodo = useCallback(async () => {
    if (!id) return;
    const response = await getTodo({ id });
    if (!response) return;
    setValue("title", response.title);
    setValue("content", response.content);
  }, [id, setValue]);

  const handleEditSubmit = handleSubmit(
    useCallback(
      async (values: z.infer<typeof schema>) => {
        if (!id) return;
        await updateTodo({
          id,
          title: values.title,
          content: values.content,
        });
        navigate(NAVIGATION_PATH.TOP);
      },
      [navigate, id]
    )
  );

  useEffect(() => {
    fetchTodo();
  }, [fetchTodo]);

  return (
    <BaseLayout title={"TodoEdit"}>
      <form className={styles.container} onSubmit={handleEditSubmit}>
        <div className={styles.area}>
          <Controller
            name="title"
            render={({ field }) => (
              <InputFormSection
                placeholder={"Title"}
                errorMessage={errors.title?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>
        <div className={styles.area}>
          <Controller
            name="content"
            render={({ field }) => (
              <TextAreaSection
                placeholder={"Content"}
                errorMessage={errors.content?.message}
                {...field}
              />
            )}
            control={control}
          />
        </div>
        <div className={styles.area}>
          <CommonButton type="submit">{"Edit Todo"}</CommonButton>
        </div>
      </form>
    </BaseLayout>
  );
};
