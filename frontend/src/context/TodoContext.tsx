import React, { createContext, useState } from 'react';
import type { ReactNode } from 'react';


interface TodoContextType {
  todos: any[];
  setTodos: React.Dispatch<React.SetStateAction<any[]>>;
}

export const TodoContext = createContext<TodoContextType | undefined>(undefined);

export const TodoProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [todos, setTodos] = useState<any[]>([]);

  return (
    <TodoContext.Provider value={{ todos, setTodos }}>
      {children}
    </TodoContext.Provider>
  );
};
