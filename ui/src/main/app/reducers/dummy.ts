export interface DummyState {
}

const initialState: DummyState = {
};

export default function(state = initialState, { type, payload }): DummyState {
    switch (type) {
        case "DUMMY":
            return {
                ...state,
                user: payload
            };
        
        default:
            return state;
    }
}