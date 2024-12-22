FROM golang

ADD ./erdis.kvo /erdis.kvo

CMD /erdis.kvo
