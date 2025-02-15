import { useMemo } from "react";
import { useForm, Controller } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { InputFormSection } from "../../molecules";
import { BaseLayout, TodoList } from "../../organisms";
import { useTodoContext } from "../../../hooks/useTodoContext";
import styles from "./style.module.css";

const schema = z.object({
  keyword: z.string(),
});

export const TodoListTemplate = () => {
  // コンテキストから状態とロジックを呼び出してコンポーネントにあてがう
  const { originTodoList, deleteTodo } = useTodoContext();

  const { control, watch } = useForm({
    resolver: zodResolver(schema),
    defaultValues: { keyword: "" },
  });
  const searchKeyword = watch("keyword");

  /* 表示用TodoList */
  const showTodoList = useMemo(() => {
    const regexp = new RegExp("^" + searchKeyword, "i");
    return originTodoList.filter((todo) => {
      // 検索キーワードに部分一致したTodoだけを一覧表示する
      return todo.title.match(regexp);
    });
    // useMemoの第二引数([originTodoList, searchKeyword])に依存して処理が実行される
    // originTodoListとsearchKeywordの値が変更される度にfilterの検索処理が実行
    // ただし結果が前回と同じならキャッシュを返却し処理は実行されない(無駄な処理を省いている)
    // 詳しくはuseMemoを調べてください。
  }, [originTodoList, searchKeyword]);

  return (
    <BaseLayout title={"TodoList"}>
      <div className={styles.container}>
        {/* Todo検索フォームエリア */}
        <div className={styles.area}>
          <Controller
            name="keyword"
            render={({ field }) => (
              <InputFormSection placeholder={"Search Keyword"} {...field} />
            )}
            control={control}
          />
        </div>
        {/* Todoリスト一覧表示 */}
        <div className={styles.area}>
          {showTodoList.length > 0 && (
            <TodoList todoList={showTodoList} handleDeleteTodo={deleteTodo} />
          )}
        </div>
      </div>
    </BaseLayout>
  );
};
