import 'package:mod_main/modules/support_roles/data/support_role_answer_model.dart';
import 'package:mod_main/core/shared_repositories/base_repository.dart';
import 'package:mod_main/core/shared_repositories/support_role_answer_repository.dart';

class MockSupportRoleAnswerRepository extends BaseRepository
    implements SupportRoleAnswerRepository {
  List<SupportRoleAnswer> getAll() {
    // Might look something like this
    // this.getConnection().query("SELECT * FROM xorgs");
    // OR if the client acts as the repository, then this repo will just mirror it
    // and do this.getConnection().getAllOrgs()

    return this._mockSupportRoleAnswers;
  }

  SupportRoleAnswer getById(String id) {
    return this._mockSupportRoleAnswers.singleWhere((_supportRoleAnswer) => _supportRoleAnswer.id == id);
  }

  // Returns a list of Orgs via a matching name
  List<SupportRoleAnswer> getByQuestionId(String questionId) {
    return this._mockSupportRoleAnswers.where((_supportRoleAnswer) => _supportRoleAnswer.refQuestionId == questionId);
  }

  List<SupportRoleAnswer> _mockSupportRoleAnswers = [
    SupportRoleAnswer(
      id: "001",
      prod: "1",
      refUserId: "001",
      refQuestionId: "001",
      answer: "3",
    ),
    SupportRoleAnswer(
      id: "002",
      prod: "1",
      refUserId: "001",
      refQuestionId: "002",
      answer: "2",
    ),
    SupportRoleAnswer(
      id: "003",
      prod: "1",
      refUserId: "001",
      refQuestionId: "003",
      answer: "3",
    ),
    SupportRoleAnswer(
      id: "004",
      prod: "1",
      refUserId: "002",
      refQuestionId: "001",
      answer: "1",
    ),
    SupportRoleAnswer(
      id: "005",
      prod: "1",
      refUserId: "003",
      refQuestionId: "004",
      answer: "2",
    ),
  ];
}
