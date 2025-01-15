import { paths } from "./apiSchema";
import { UnionToIntersection, Get } from "type-fest";
import axios, { AxiosResponse, AxiosRequestConfig } from "axios";

export type UrlPaths = keyof paths;

export type HttpMethods = Exclude<keyof UnionToIntersection<paths[keyof paths]>, symbol>;

export type RequestParameters<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.parameters.query`>;

export type RequestPathParam<
  Path extends UrlPaths,
  Method extends HttpMethods
> = Get<paths, `${Path}.${Method}.parameters.path`>;

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

export type AxiosConfigWrapper<Path extends UrlPaths, Method extends HttpMethods> = AxiosRequestConfig  & {
	url: Path;
	method: Method & HttpMethodsFilteredByPath<Path>;
	params?: RequestParameters<Path, Method>;
	data?: RequestData<Path, Method>;
  pathParam?: RequestPathParam<Path, Method>;
}

export async function request<Path extends UrlPaths, Method extends HttpMethods>(config: AxiosConfigWrapper<Path, Method>) {
  const requestConfig: AxiosRequestConfig = {
    ...config,
  }

  if (config.pathParam) {
    for (const prop in config.pathParam) {
      const value = config.pathParam[prop]
      requestConfig.url = requestConfig.url?.replace(`{${prop}}`, `${value}`)
    }
  }
  return await axios.request<
    ResponseData<Path, Method>,
    AxiosResponse<ResponseData<Path, Method>>,
    AxiosConfigWrapper<Path, Method>["data"]
  >(requestConfig);
}
