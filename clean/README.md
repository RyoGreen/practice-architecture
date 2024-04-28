# clean architecture

## overall
controller -> usecase -> entity  
    |          |  
     repository

## entity
- staffのビジネスモデルを作成
- 例としてビジネスルールとしてvalidationを作成

## usecase
- controllerから渡されるデータをentityのデータに変換
- repository層にてCRUDを実施する
- validationを実施
- controllertに返すデータに変換する

## controller
- controller.goでhttp通信を前提として、usecaseに渡すデータを作成する
- usecaseから取得したデータをresponseに返すようにする

## repositoty
- データを扱うツールの具体的な実装をrepository内で実装する
- ここでは例としてmap(非現実的)とmysqlにて実装をしている
- usecaes interface -> usecase -> repository interface -> repository
