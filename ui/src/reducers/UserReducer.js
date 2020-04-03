const UserReducer = (state, action) => {
  switch (action.type) {
    case 'LOGIN':
      return {
        ...state,
        token: action.token,
        login: true,
      };
    case 'LOGOUT':
      return {
        ...state,
        token: null,
        login: false,
      };
    case 'ERROR':
      return {
        ...state,
        error: action.error,
      };
    default:
      return state;
  }
};
export default UserReducer;
