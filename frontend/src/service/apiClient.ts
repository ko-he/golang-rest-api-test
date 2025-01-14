import { paths } from "./apiSchema";
import { UnionToIntersection, Get } from "type-fest";
import axios, { AxiosResponse, AxiosRequestConfig } from "axios";

export type UrlPaths = keyof paths;

export type HttpMethods = Exclude<keyof UnionToIntersection<paths[keyof paths]>, symbol>;

export type RequestParameters<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.parameters.query`>;

export type HttpMethodsFilteredByPath<Path extends UrlPaths> = HttpMethods &
  keyof UnionToIntersection<paths[Path]>;

export type RequestData<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.requestBody.content.application/json`>;

export type ResponseData<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.responses.200.content.application/json`>;

export type AxiosConfigWrapper< Path extends UrlPaths, Method extends HttpMethods>  = AxiosRequestConfig  & {
	url: Path;
	method: Method & HttpMethodsFilteredByPath<Path>;
	params?: RequestParameters<Path, Method>;
	data?: RequestData<Path, Method>;
}

export async function request<Path extends UrlPaths, Method extends HttpMethods>(config: AxiosConfigWrapper<Path, Method>) {
  return await axios.request<
    ResponseData<Path, Method>,
    AxiosResponse<ResponseData<Path, Method>>,
    AxiosConfigWrapper<Path, Method>["data"]
  >(config);
}
