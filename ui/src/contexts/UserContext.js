import React, {createContext, useReducer} from "react";
import UserReducer from '../reducers/UserReducer';

const initialState = {
  token: null,
  error: false,
  login: true
};

const UserStore = ({children}) => {
  const [state, dispatch] = useReducer(UserReducer, initialState);
  return (
    <UserContext.Provider value={[state, dispatch]}>
    {children}
    </UserContext.Provider>

  );

};

export const UserContext = createContext(initialState);
export default UserStore;
