export class ServiceURL {
    public static HOST: string = '/api';

    public static loginURL: string =  `${ServiceURL.HOST}/auth/login/`;
    public static logoutURL: string = `${ServiceURL.HOST}/auth/logout/`;

    public static todoURL: string = `${ServiceURL.HOST}/todos/`;
    public static todoMarkDoneURL: string = `${ServiceURL.HOST}/todos/done/`;
}
