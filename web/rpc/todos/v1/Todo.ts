//
// GENERATED CODE -- DO NOT EDIT!
//
// source: todos/v1/Todo.proto
//
export enum TodoState {
  UNKNOWN = 0,
  ARCHIVED = 1,
  ACTIVE = 2,
}
export interface Todo {
  id?: string;
  title?: string;
  state?: TodoState;
}