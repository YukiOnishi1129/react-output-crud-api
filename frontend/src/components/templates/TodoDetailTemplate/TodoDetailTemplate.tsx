import { useParams } from "react-router";
import { useCallback, useState, useEffect } from "react";
import { BaseLayout } from "../../organisms";
import { InputForm, TextArea } from "../../atoms";
import { getTodo } from "../../../apis/todo";
import { TodoType } from "../../../types/Todo";
import styles from "./style.module.css";

export const TodoDetailTemplate = () => {
  const { id } = useParams();

  const [todo, setTodo] = useState<TodoType | null>(null);

  const fetchTodo = useCallback(async () => {
    if (!id) return;
    const response = await getTodo({ id });
    if (!response) return;

    setTodo(response);
  }, [id]);

  useEffect(() => {
    fetchTodo();
  }, [fetchTodo]);

  return (
    <BaseLayout title={"TodoDetail"}>
      <div></div>
      {!!todo && (
        <div className={styles.container}>
          <div className={styles.area}>
            <InputForm disabled value={todo.title} placeholder={"Title"} />
          </div>
          <div className={styles.area}>
            <TextArea disabled value={todo.content} placeholder={"Content"} />
          </div>
        </div>
      )}
    </BaseLayout>
  );
};
