package public

// func (obj *testObject) TestGetOrderSuccess() {
// 	obj.MockService.EXPECT().OrderList(gomock.Any()).Return([]*dto.OrderDTO{
// 		{
// 			OrderStores: []dto.OrderStoreDTO{
// 				{
// 					OrderStoreItems: []entity.OrderStoreItem{
// 						{},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			OrderStores: []dto.OrderStoreDTO{
// 				{
// 					OrderStoreItems: []entity.OrderStoreItem{
// 						{},
// 					},
// 				},
// 			},
// 		},
// 	}, nil)
// 	response := obj.module.GetOrder(obj.writer, obj.request)
// 	require.NotNil(obj.T(), response)
// }

// func (obj *testObject) TestGetOrderFailed() {
// 	obj.MockService.EXPECT().OrderList(gomock.Any()).Return([]*dto.OrderDTO{}, errors.New("something Bad Happen"))
// 	response := obj.module.GetOrder(obj.writer, obj.request)
// 	require.NotNil(obj.T(), response)
// }
