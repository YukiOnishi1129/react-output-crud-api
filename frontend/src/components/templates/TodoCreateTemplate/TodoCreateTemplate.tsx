import { useCallback } from "react";
import { Controller, useForm } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "react-router";
import { NAVIGATION_PATH } from "../../../constants/navigation";

import { BaseLayout } from "../../organisms";
import { InputFormSection, TextAreaSection } from "../../molecules";
import { CommonButton } from "../../atoms";

import { createTodo } from "../../../apis/todo";

import styles from "./style.module.css";

const schema = z.object({
  title: z
    .string()
    .min(1, "タイトルは必須です。")
    .max(10, "10文字以内で入力してください。"),
  content: z.string().optional(),
});

export const TodoCreateTemplate = () => {
  const navigate = useNavigate();

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: { title: "", content: undefined },
  });

  const handleAddSubmit = handleSubmit(
    useCallback(
      async (values: z.infer<typeof schema>) => {
        await createTodo({
          title: values.title,
          content: values.content,
        });
        navigate(NAVIGATION_PATH.TOP);
      },
      [navigate]
    )
  );

  return (
    <BaseLayout title={"Create Todo"}>
      <div></div>
      <form className={styles.container} onSubmit={handleAddSubmit}>
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
          <CommonButton type="submit">{"Create Todo"}</CommonButton>
        </div>
      </form>
    </BaseLayout>
  );
};
