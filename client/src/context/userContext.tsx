// @ts-nocheck
import { createContext } from 'react';
import React, { useReducer, useContext, ReactNode, Dispatch } from 'react'; // Import Dispatch from react
import User from '../types/user';
import Plan from '../types/plan';

// Define the Action types
type Action = { type: 'UPDATE_USER'; payload: User };

// Define the ContextType interface
interface ContextType {
  user: User;
  dispatch: Dispatch<Action>; // Use Dispatch<Action> for dispatch function
}

// Define the initial plan
const initialPlan: Plan = {
  planName: '',
  uploadLimitSizeKb: 0, // Corrected property name
  dailyUploadLimit: 0,
  cost: 0,
};

// Define the initial user state
const initialState: User = {
  email: '',
  username: '',
  firstName: '',
  lastName: '',
  pictureURI: '',
  plan: initialPlan,
};

// Define the reducer function
const reducer = (state: User, action: Action): User => {
  switch (action.type) {
    case 'UPDATE_USER':
      return { ...state, ...action.payload };
    default:
      return state;
  }
};

// Create the UserContext
export const UserContext = createContext<ContextType | undefined>(undefined);

// Define the UserProvider component
const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [user, dispatch] = useReducer(reducer, initialState);

  return (
    <UserContext.Provider value={{ user, dispatch }}>
      {children}
    </UserContext.Provider>
  );
};

// Define the custom hook useUserContext
const useUserContext = (): ContextType => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error('useUserContext must be used within a UserProvider');
  }
  return context;
};

export { UserProvider, useUserContext };
