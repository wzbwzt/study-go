## 什么是DI(Dependency Inject，简称DI)
> 把有依赖关系的类放到容器中，解析出这些类的实例，就是依赖注入。

### go中有两个包实现了DI,dig和wire
dig 主要使用反射，不够高效
wire 更加高效好用