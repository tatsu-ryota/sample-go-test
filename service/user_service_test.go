package service

import (
	"errors"
	"testing"

	"tatsu-ryota/sample-go-test/mocks"
	"tatsu-ryota/sample-go-test/repository"

	"github.com/golang/mock/gomock"
)

// TestGetUserName: GetUserName の動作をテストする
func TestGetUserName(t *testing.T) {
	// gomock のコントローラーを作成（モックの動作を管理）
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // テスト終了後にリソースを解放

	// モックのリポジトリを作成（本物のデータベースを使わないため）
	mockRepo := mocks.NewMockUserRepository(ctrl)

	// テスト対象の UserService を作成（モックを利用）
	service := NewUserService(mockRepo)

	// テストケース（複数のパターンをリスト化）
	tests := []struct {
		name    string           // テストの名前（説明用）
		userID  int              // テスト対象のユーザーID
		mockOut *repository.User // モックが返すユーザー情報
		mockErr error            // モックが返すエラー
		want    string           // 期待する戻り値
		wantErr bool             // 期待するエラーの有無
	}{
		{
			name:    "正常系: ユーザー取得成功",                        // 期待通りの結果が返るケース
			userID:  1,                                      // 入力（ユーザーID）
			mockOut: &repository.User{ID: 1, Name: "Alice"}, // モックの戻り値
			mockErr: nil,                                    // エラーなし
			want:    "Alice",                                // 期待する結果
			wantErr: false,                                  // エラーなし
		},
		{
			name:    "異常系: ユーザーが見つからない",           // エラーが発生するケース
			userID:  1,                            // 入力（ユーザーID）
			mockOut: nil,                          // ユーザー情報なし
			mockErr: errors.New("user not found"), // エラー発生
			want:    "",                           // 期待する結果（エラー時は空文字）
			wantErr: true,                         // エラー発生を期待
		},
	}

	// 各テストケースを順番に実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // テストケースごとに独立したサブテストを実行
			// モックの設定: GetUser() が呼ばれたら mockOut, mockErr を返す
			mockRepo.EXPECT().GetUser(tt.userID).Return(tt.mockOut, tt.mockErr)
			// GetUserの引数の指定がなければ以下の書き方でも可
			// mockRepo.EXPECT().GetUser(gomock.Any()).Return(tt.mockOut, tt.mockErr)

			// テスト対象の関数を実行
			got, err := service.GetUserName(tt.userID)

			// 期待するエラーと実際のエラーを比較
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserName(%d) error = %v; wantErr %v", tt.userID, err, tt.wantErr)
			}

			// 期待する値と実際の値を比較
			if got != tt.want {
				t.Errorf("GetUserName(%d) = %q; want %q", tt.userID, got, tt.want)
			}
		})
	}
}
