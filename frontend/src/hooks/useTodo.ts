import { useState, useCallback } from "react";
import { INIT_TODO_LIST, INIT_UNIQUE_ID } from "../constants/data";

export const useTodo = () => {
  /* todolist */
  const [originTodoList, setOriginTodoList] = useState(INIT_TODO_LIST);
  /* todo採番ID */
  const [uniqueId, setUniqueId] = useState(INIT_UNIQUE_ID);

  /* actions */
  /**
   * Todo新規登録処理
   */
  const addTodo = useCallback(
    (title: string, content?: string) => {
      const nextUniqueId = uniqueId + 1;
      const newTodo = [
        ...originTodoList,
        {
          id: nextUniqueId,
          title: title,
          content: content,
        },
      ];
      // todolistを更新
      setOriginTodoList(newTodo);
      // 採番IDを更新
      setUniqueId(nextUniqueId);
    },
    [originTodoList, uniqueId]
  );

  /**
   * Todo更新処理
   */
  const updateTodo = useCallback(
    (id: number, title: string, content?: string) => {
      const updatedTodoList = originTodoList.map((todo) => {
        if (id === todo.id) {
          return {
            id: todo.id,
            title: title,
            content: content,
          };
        }

        return todo;
      });
      setOriginTodoList(updatedTodoList);
    },
    [originTodoList]
  );

  /**
   * Todo削除処理
   */
  const deleteTodo = useCallback(
    (targetId: number, targetTitle: string) => {
      if (window.confirm(`「${targetTitle}」のtodoを削除しますか？`)) {
        // 削除するid以外のtodoリストを再編集
        // filterを用いた方法
        const newTodoList = originTodoList.filter(
          (todo) => todo.id !== targetId
        );

        // 削除するTodoの配列番号を取り出してspliceで削除する方法もある
        // const newTodoList = [...todoList];
        // const deleteIndex = newTodoList.findIndex((todo) => todo.id === targetId);
        // newTodoList.splice(deleteIndex, 1);

        // todoを削除したtodo listで更新
        setOriginTodoList(newTodoList);
      }
    },
    [originTodoList]
  );

  return {
    originTodoList,
    addTodo,
    updateTodo,
    deleteTodo,
  };
};
