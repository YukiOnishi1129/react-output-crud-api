import { BaseLayout } from "../../organisms";

export const TodoEditTemplate = () => {
  // const { originTodoList, updateTodo } = useTodoContext();

  // const { todo, control, errors, handleEditSubmit } = useTodoEditTemplate({
  //   originTodoList,
  //   updateTodo,
  // });

  return (
    <BaseLayout title={"TodoEdit"}>
      <div></div>
      {/* {!!todo && (
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
      )} */}
    </BaseLayout>
  );
};
