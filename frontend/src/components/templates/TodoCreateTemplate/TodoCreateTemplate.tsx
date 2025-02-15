import { useTodoContext } from "../../../hooks/useTodoContext";
import { BaseLayout } from "../../organisms";
import { CommonButton, InputForm, TextArea } from "../../atoms";
import { useTodoCreateTemplate } from "./useTodoCreateTemplate";
import styles from "./style.module.css";

export const TodoCreateTemplate = () => {
  const { addTodo } = useTodoContext();
  const {
    inputTitle,
    inputContent,
    handleChangeTitle,
    handleChangeContent,
    handleCreateTodo,
  } = useTodoCreateTemplate({ addTodo });

  return (
    <BaseLayout title={"Create Todo"}>
      <form className={styles.container} onSubmit={handleCreateTodo}>
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
          <CommonButton type="submit">{"Create Todo"}</CommonButton>
        </div>
      </form>
    </BaseLayout>
  );
};
