import { useTodoEditTemplate } from "./useTodoEditTemplate";
import { useTodoContext } from "../../../hooks/useTodoContext";
import { BaseLayout } from "../../organisms";
import { CommonButton, TextArea, InputForm } from "../../atoms";
import styles from "./style.module.css";

export const TodoEditTemplate = () => {
  const { originTodoList, updateTodo } = useTodoContext();

  const {
    todo,
    inputTitle,
    inputContent,
    handleChangeTitle,
    handleChangeContent,
    handleUpdateTodo,
  } = useTodoEditTemplate({ originTodoList, updateTodo });

  return (
    <BaseLayout title={"TodoEdit"}>
      {!!todo && (
        <form className={styles.container} onSubmit={handleUpdateTodo}>
          <div className={styles.area}>
            <InputForm
              value={inputTitle}
              placeholder={"Title"}
              onChange={handleChangeTitle}
            />
          </div>
          <div className={styles.area}>
            <TextArea
              value={inputContent}
              placeholder={"Content"}
              onChange={handleChangeContent}
            />
          </div>
          <div className={styles.area}>
            <CommonButton type="submit">{"Edit Todo"}</CommonButton>
          </div>
        </form>
      )}
    </BaseLayout>
  );
};
