export const dummy = () => {
    return async (dispatch) => {
        dispatch({
            type: "DUMMY",
            payload: {}
        });
    };
};