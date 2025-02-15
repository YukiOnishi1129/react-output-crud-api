import { useMemo, useState, useCallback, ChangeEvent } from "react";
import { TodoType } from "../../../types/Todo";

type UseTodoListTemplateParams = {
  originTodoList: Array<TodoType>;
};

export const useTodoListTemplate = ({
  originTodoList,
}: UseTodoListTemplateParams) => {
  /* 検索キーワード */
  const [searchKeyword, setSearchKeyword] = useState("");
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

  /**
   * 検索キーワード更新処理
   */
  const handleChangeSearchKeyword = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => setSearchKeyword(e.target.value),
    []
  );

  return { searchKeyword, showTodoList, handleChangeSearchKeyword };
};
