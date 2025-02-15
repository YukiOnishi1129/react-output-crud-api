import { useMemo, useState, useEffect, useCallback } from "react";
import { useForm, Controller } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { InputFormSection } from "../../molecules";
import { BaseLayout, TodoList } from "../../organisms";

import { getTodos, deleteTodo } from "../../../apis/todo";
import { TodoType } from "../../../types/Todo";

import styles from "./style.module.css";

const schema = z.object({
  keyword: z.string(),
});

export const TodoListTemplate = () => {
  const [todoList, setTodoList] = useState<Array<TodoType>>([]);

  const { control, watch } = useForm({
    resolver: zodResolver(schema),
    defaultValues: { keyword: "" },
  });
  const searchKeyword = watch("keyword");

  /* 表示用TodoList */
  const showTodoList = useMemo(() => {
    const regexp = new RegExp("^" + searchKeyword, "i");
    return todoList.filter((todo) => {
      // 検索キーワードに部分一致したTodoだけを一覧表示する
      return todo.title.match(regexp);
    });
    // useMemoの第二引数([originTodoList, searchKeyword])に依存して処理が実行される
    // originTodoListとsearchKeywordの値が変更される度にfilterの検索処理が実行
    // ただし結果が前回と同じならキャッシュを返却し処理は実行されない(無駄な処理を省いている)
    // 詳しくはuseMemoを調べてください。
  }, [todoList, searchKeyword]);

  const fetchTodoList = useCallback(async () => {
    const response = await getTodos();
    if (!response) return;

    setTodoList(response.todos);
  }, []);

  const handleDeleteTodo = useCallback((id: string, title: string) => {
    if (window.confirm(`「${title}」のtodoを削除しますか？`)) {
      deleteTodo({ id });
    }
  }, []);

  useEffect(() => {
    fetchTodoList();
  }, [fetchTodoList]);

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
            <TodoList todoList={showTodoList} onDeleteTodo={handleDeleteTodo} />
          )}
        </div>
      </div>
    </BaseLayout>
  );
};
