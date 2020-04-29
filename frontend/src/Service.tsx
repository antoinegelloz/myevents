export interface ServiceInit {
    status: 'init';
}
export interface ServiceLoading {
    status: 'loading';
}
export interface ServiceReady<T> {
    status: 'ready';
    payload: T;
}
export interface ServiceSaving {
    status: 'saving';
}
export interface ServiceDone {
    status: 'done';
}
export interface ServiceError {
    status: 'error';
    error: Error;
}
export type Service<T> =
    | ServiceInit
    | ServiceLoading
    | ServiceReady<T>
    | ServiceSaving
    | ServiceDone
    | ServiceError;